#!/bin/bash
set -euxo pipefail

# Colorful output.
function greenprint {
    echo -e "\033[1;32m${1}\033[0m"
}

function retry {
    local count=0
    local retries=5
    until "$@"; do
        exit=$?
        count=$((count + 1))
        if [[ $count -lt $retries ]]; then
            echo "Retrying command..."
            sleep 1
        else
            echo "Command failed after ${retries} retries. Giving up."
            return $exit
        fi
    done
    return 0
}

# Get OS details.
source /etc/os-release

if [[ -n "${RHN_REGISTRATION_SCRIPT:-}" ]] && ! sudo subscription-manager status; then
    greenprint "Registering RHEL"
    sudo chmod +x "$RHN_REGISTRATION_SCRIPT"
    sudo "$RHN_REGISTRATION_SCRIPT"
fi

greenprint "Restarting systemd to work around some Fedora issues in cloud images"
sudo systemctl restart systemd-journald

greenprint "Removing Fedora's modular repositories to speed up dnf"
sudo rm -f /etc/yum.repos.d/fedora*modular*

greenprint "Enabling fastestmirror and disabling weak dependencies to speed up dnf even more 🏎️"
echo -e "fastestmirror=1\ninstall_weak_deps=0" | sudo tee -a /etc/dnf/dnf.conf

# Ensure we are using the latest dnf since early revisions of Fedora 31 had
# some dnf repo priority bugs like BZ 1733582.
# NOTE(mhayden): We can exclude kernel updates here to save time with dracut
# and module updates. The system will not be rebooted in CI anyway, so a
# kernel update is not needed.
if [[ $ID == fedora ]]; then
    greenprint "Upgrading system to fix dnf issues"
    sudo dnf -y upgrade --exclude kernel --exclude kernel-core
fi

greenprint "Adding osbuild team ssh keys"
cat schutzbot/team_ssh_keys.txt | tee -a ~/.ssh/authorized_keys > /dev/null

greenprint "Setting up a dnf repository for the RPMs we built via mock"
sudo cp osbuild-mock.repo /etc/yum.repos.d/osbuild-mock.repo
sudo dnf repository-packages osbuild-mock list

if [[ $ID == rhel ]]; then
    greenprint "Setting up EPEL repository"
    # we need this for ansible and koji
    sudo dnf install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm
fi

greenprint "Installing the Image Builder packages"
# Note: installing only -tests to catch missing dependencies
retry sudo dnf -y install osbuild-composer-tests

greenprint "Setting up a directory to hold repository overrides for weldr"
sudo mkdir -p /etc/osbuild-composer/repositories

# Copy Fedora rpmrepo snapshots for use in weldr tests. RHEL's are usually more
# stable, and not available publically from rpmrepo.
sudo cp schutzbot/repositories/fedora-*.json /etc/osbuild-composer/repositories/

greenprint "Copying repository configuration for tests"
sudo mkdir -p /etc/tests/osbuild-composer/repositories
sudo cp schutzbot/repositories/*.json /etc/tests/osbuild-composer/repositories/

greenprint "Provisioning the services"
./schutzbot/provision.sh

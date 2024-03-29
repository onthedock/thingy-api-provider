#!/usr/bin/env bash
sudo apt-get install make --yes
grep --quiet --fixed-strings --line-regexp 'source .devcontainer/git-completion.bash' ~/.bashrc || echo 'source .devcontainer/git-completion.bash' >> ~/.bashrc

echo "Git-delta (AMD64)" 
culr -JLO https://github.com/dandavison/delta/releases/download/0.17.0/git-delta_0.17.0_amd64.deb
sudo dpkg -i git-delta_0.17.0_amd64.deb

# Install Terraform
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform

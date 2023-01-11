# Scripts

This folder contains a sequence of helper scripts for creating an furya on testnet and automatic delegation. 

1. **gov.sh** submits a gov.json governance proposal, votes in favor of it and then queries the created furya.
2. **delegate.sh** delegates to the previously create furya and queries the modified furya.
3. **rewards.sh** claims available rewards and retrieves information about the process

> Note that these scripts must be executed in the specified order since they have dependencies on each other.

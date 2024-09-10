# Technical Specification for Aruba Cloud Provider Command Line Interface (CLI)

## Project Name: Aruba Cloud Provider CLI (acloud)

### Objective

Develop a command-line interface (CLI) for interacting with the Aruba cloud provider infrastructure and services. The CLI should leverage a set of public APIs and allow users to manage their cloud infrastructure, services, and resources efficiently from a terminal or automated scripts.

## Key Requirements

### 1. Architecture & Design

#### Programming Language
- The CLI will be implemented in a language suitable for command-line tool development, such as Go.

#### API Integration
- The CLI will integrate with the Aruba Cloud provider public REST APIs.
- API calls will be authenticated using OAuth2.0 client credentials grant type.

#### Platform Compatibility
- Cross-platform support (Linux, Windows, macOS).
- Built-in package management for easy installation using package managers like `pip`, `brew`, and `apt`.

### 2. Core Functionality

The CLI should implement the following key features, mapped to cloud services via API endpoints:

#### Infrastructure Management
- `create`: Create cloud resources (e.g., VM instances, Kubernetes cluster).
- `delete`: Delete cloud resources.
- `list`: List resources (e.g., VMs, volumes, networks).
- `describe`: Fetch details of a specific resource.

### 2.1 Details

#### Authentication
- `login`: Authenticate users with OAuth2.0 Client Credentials Grant Type.
  - The web portal would be used to generate API clients with the related secret value.
- `logout`: Terminate the session.

#### Management

##### Project
- `create-project`: Create a new Project.
- `describe-project`: Get Project Details.
- `update-project`: Update a Project.
- `delete-project`: Delete a Project.
- `list-projects`: List Projects.
- `list-resources`: List Project Resources.

##### Container

###### KaaS
- `create-kaas`: Create a new KaaS.
- `describe-kaas`: Get KaaS Details.
- `update-kaas`: Update a KaaS.
- `delete-kaas`: Delete a KaaS.
- `list-kaas`: List available KaaS.

##### Compute

###### CloudServer
- `create-cloudserver`: Create a new Cloud Server.
- `describe-cloudserver`: Get Cloud Server Detail.
- `update-cloudserver`: Update Cloud Server.
- `delete-cloudserver`: Delete a Cloud Server.
- `list-cloudservers`: List available Cloud Servers.

###### CloudServerSnapshot
- `create-cloudserversnapshot`: Create a new Cloud Server Snapshot.
- `describe-cloudserversnapshot`: Get Cloud Server Snapshot details.
- `update-cloudserversnapshot`: Update a Cloud Server Snapshot.
- `delete-cloudserversnapshot`: Delete a Cloud Server Snapshot.
- `list-cloudserversnapshots`: List Cloud Server Snapshots.

###### KeyPair
- `create-keypair`: Create a new KeyPair.
- `describe-keypair`: Get KeyPair details.
- `update-keypair`: Update KeyPair.
- `delete-keypair`: Delete a KeyPair.
- `list-keypairs`: List KeyPairs.

##### Network

###### VPC
- `create-vpc`: Create a new virtual network.
- `describe-vpc`: Get virtual network details.
- `update-vpc`: Update virtual network.
- `delete-vpc`: Delete a virtual network.
- `list-vpcs`: List available virtual networks.

###### Subnet
- `create-subnet`: Create a new Subnet in a virtual network.
- `describe-subnet`: Get Subnet details.
- `update-subnet`: Update Subnet.
- `delete-subnet`: Delete a Subnet.
- `list-subnets`: List available subnets in a virtual network.

###### SecurityGroup
- `create-securitygroup`: Create a new Security Group.
- `describe-securitygroup`: Get Security Group details.
- `update-securitygroup`: Update Security Group.
- `delete-securitygroup`: Delete a Security Group.
- `list-securitygroups`: List available Security Groups in a virtual network.

###### SecurityRule
- `create-securityrule`: Create a new Security Rule in a Security Group of a virtual network.
- `describe-securityrule`: Get Security Rule details.
- `update-securityrule`: Update a Security Rule.
- `delete-securityrule`: Delete a Security Rule.
- `list-securityrules`: List available Security Rules of a Security Group.

###### ElasticIp
- `create-elasticip`: Create a new Elastic IP.
- `describe-elasticip`: Get Elastic IP details.
- `update-elasticip`: Update an Elastic IP.
- `delete-elasticip`: Delete an Elastic IP.
- `list-elasticips`: List available Elastic IPs.

###### Load Balancer
- `describe-loadbalancer`: Get Load Balancer details.
- `list-loadbalancers`: List available Load Balancers.

##### Storage

###### BlockStorage
- `create-blockstorage`: Create storage volumes.
- `describe-blockstorage`: Get Storage Volume details.
- `update-blockstorage`: Update Storage Volume.
- `delete-blockstorage`: Delete storage volumes.
- `attach-blockstorage`: Attach a volume to a VM.
- `list-blockstorages`: List storage volumes.

##### Monitor

###### Metric
- `statistics`: Retrieve metrics for cloud services.

###### Audit
- `logs`: Retrieve audit logs for cloud services.


### Examples

```bash
# Authenticate with the cloud provider
$ acloud login --client-id YOUR_API_KEY --client-secret YOUR_SECRET_VALUE
```

```bash
# Example Create Project
$ acloud management create-project <....details.....>
```

```bash
# Example Create KaaS
$ acloud container create-kaas <....details.....>
```

### 3. CLI User Experience

#### Command Structure
- Use subcommands for different services (e.g., `acloud compute list`, `acloud storage create`).
- Each subcommand should follow a consistent pattern: command + resource + action (e.g., `create-vm`, `list-volumes`, `delete-firewall-rule`).

#### Output Formats
- Support multiple output formats, such as JSON, YAML, and table for easy parsing in scripts.
- Example: `acloud compute list --output json`.

#### Command Options and Flags
- Include command-line flags for filtering, pagination, and output customization.
- Example: `acloud compute list --region us-east --status running`.

#### Interactive Mode
- The CLI should offer an interactive mode (e.g., `acloud interactive`), allowing users to execute multiple commands without relaunching the CLI.

### 4. Error Handling & Logging

#### Error Handling
- Provide detailed error messages for failed operations, indicating the HTTP status code, API response, and possible resolutions.

#### Logging
- Option to enable verbose logging (`--verbose`) for debugging.
- Write logs to a default log file in the user’s home directory or a configurable path.

### 5. Security

#### Authentication Tokens
- Store authentication tokens securely in the system’s credential store (e.g., Keychain for macOS, Windows Credential Manager, etc.).

#### TLS Encryption
- Ensure all API calls are made over HTTPS, enforcing TLS 1.2 or higher.

### 6. Performance Considerations

#### Parallel Requests
- Use multithreading for certain operations, such as listing or deleting resources in bulk, to minimize latency.

#### Caching
- Cache API responses temporarily (configurable) to reduce redundant API calls (e.g., `acloud describe-vm` may use cached data if called multiple times within a short period).

### 7. Modular and Extensible Design

#### Plugin Support
- Implement a plugin architecture that allows the community to extend the CLI with new features or support for additional services.

#### Configurable Defaults
- Allow users to configure default parameters (e.g., default region, output format) using a configuration file (`~/.acloud/config.yaml`).

### 8. Testing & Quality Assurance

#### Unit Testing
- Ensure unit test coverage for all core functionalities, particularly API interactions, response parsing, and error handling.

#### Integration Testing
- Integration tests to validate API calls and data manipulation (mock responses where necessary).

### 9. Documentation & Help

#### Help Command
- Provide detailed inline help (`acloud --help`) with examples for each command and subcommand.

#### Man Pages
- Offer comprehensive man pages for offline usage.

#### Online Documentation
- Host detailed documentation and examples on the cloud provider’s website or a dedicated GitHub repository.

### 10. Release & Versioning

#### Versioning
- Follow semantic versioning (SemVer) for releases (e.g., v1.0.0, v1.1.0).

#### Distribution
- Distribute via package manageracloud logins (e.g., `pip`, `brew`, `apt`).
- Option for standalone binaries to download directly.

### 11. CLI Use Example

```bash
# Authenticate with the Aruba Cloud Provider
acloud login --client-id YOUR_API_KEY --client-secret YOUR_SECRET_VALUE

# List all virtual machines in the ITBG-Bergamo region
acloud compute list-cloudservers --region ITBG-Bergamo

# Create a new VM instance
acloud compute create-cloudserver --name my-vm --network-id $id --flavor-id CSO8A16 --instance-type t2.micro --region us-west

# Attach a volume to the VM
acloud storage attach-blockstorage --vm-id $id --volume-id $id

# List all LoadBalancers in the ITBG-Bergamo region
acloud network list-loadbalancers --region ITBG-Bergamo

```

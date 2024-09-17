# Using Helmfile for Managing Helm Chart and Its Values File with Appropreate Mapping. 

### Helmfile Deployment

The main goal of using Helmfile in this project is to manage the `values.yml` file properly by mapping appropriate values for different environments and releases. To achieve this, the `values.yml` file is intentionally split across multiple files to ensure better organization and flexibility. This approach allows for easy handling of environment-specific values, secrets, and other configurations required for deployment.

#### Helmfile Structure:

- **Purpose:**  
  Helmfile is used to simplify and centralize the management of Helm deployments, particularly when working with complex environments. By splitting the values into different files, it allows for:
  - Easier management of environment-specific values.
  - Clear separation of configurations, such as secrets, for better security practices.
  - Scalability for adding or modifying values for different environments or clusters without affecting the entire configuration.

- **Split Values:**
  The values used in the Helm chart are split across multiple files for different environments for better management and flexibility. For each environment, the following value files are included:
  - `values/app/{{ $env_name }}/values.yml`: Main values for the app.
  - `values/app/{{ $env_name }}/secrets/secrets.yml`: Secrets specific to the environment.
  - `values/app/{{ $env_name }}/val1/val1.yml`: Additional custom values.
  - `values/app/{{ $env_name }}/val2/val2.yml`: Another custom value file for finer control.

  This split helps maintain a clear structure while handling different aspects of configuration across environments.

#### Key Features:

- **Release Name:** Set dynamically using the `RELEASE_NAME` environment variable (default: `golang-app`).
- **Namespace:** Set via the `NAMESPACE` environment variable (default: `default`).
- **Environment:** Managed via the `ENV_NAME` environment variable (default: `dev`).

- **Image Repository and Tag:**
  - `repository`: Set via the `repository` environment variable.
  - `tag`: Set via the `tag` environment variable.

  These environment variables ensure that the correct Docker image and tag are used for each deployment.

#### Setting Up and Using Helmfile:

1. Set the necessary environment variables for your deployment:

```bash
export NAMESPACE="my-namespace" //default value is default namespace
export RELEASE_NAME="my-release" //default is golang-app
export ENV_NAME="dev" //default is dev
export repository="" //required value
export tag="" //required value

helmfile sync //deploy command
```

2. The `values.yml` and other split files for the selected environment will be automatically used during deployment.

#### GitHub Actions Pipeline

This project uses a GitHub Actions pipeline for deploying Helmfile configurations. The pipeline supports two environments: `dev` and `prod`.

##### Key Steps in the Pipeline:

1. **Checkout Code:**
   Pulls the repository code into the GitHub Actions runner.

2. **Install Dependencies:**
   Installs Helm, Helmfile, and Helm Diff plugin on the runner.

   `For the self-hosted runner, Helmfile and Helm Diff plugin are installed as follows`:
   
   ```bash
   # Install Helm
   curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
   sudo apt-get install apt-transport-https --yes
   echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
   sudo apt-get update
   sudo apt-get install helm

   # Install Helm Diff Plugin
   helm plugin add https://github.com/databus23/helm-diff

   # Install Helmfile
   HELMFILE_VERSION="0.159.0"
   curl -Lo helmfile.tar.gz https://github.com/helmfile/helmfile/releases/download/v${HELMFILE_VERSION}/helmfile_${HELMFILE_VERSION}_linux_amd64.tar.gz
   tar xzf helmfile.tar.gz -C /tmp
   sudo mv /tmp/helmfile /usr/local/bin/
   chmod +x /usr/local/bin/helmfile
   rm helmfile.tar.gz
   ```

3. **Set up Local Test Cluster using Kind:**
   A local test cluster is created using Kind (`kind create cluster --name test-cluster`), and Helmfile deploys to this local cluster.

4. **Test Deployment:**
   The pipeline runs the following Helmfile commands:
   
   - **Template:** Renders the Kubernetes manifests without applying them.
     ```bash
     helmfile --kube-context kind-test-cluster template
     ```
   
   - **Diff:** Compares the rendered manifests with the currently running release.
     ```bash
     helmfile --kube-context kind-test-cluster diff
     ```

   - **Sync (Deployment):** (commented in the pipeline but can be enabled)
     ```bash
     helmfile --kube-context kind-test-cluster sync
     ```

---

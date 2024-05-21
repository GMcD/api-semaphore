# API Semaphore

Hello

## Terraform Infrastructure Deployment

```mermaid
sequenceDiagram
create participant Terraform.State
0.Bucket->>Terraform.State: Remote State Store
1.State->>Terraform.State: Create Backend
```

## Github Workflow Action Sequence

The diagram actors are described as

- Commit.\<Branch\> - a commit to that branch
- Action.\<Branch\> - a workflow action to be taken on that commit
- Google.Cloud - gcloud client actions

```mermaid
sequenceDiagram
Commit.Test->>Action.Test: Can you test my commit?
Action.Test-->>Commit.Test: Running that now..
Action.Test-xCommit.Test: Failed, fix tests...
create participant Commit.Main
Action.Test->>Commit.Main: Done, all good!
Action.Test->>Commit.Main: Committing to Main!
create participant Action.Main
Commit.Main->>Action.Main: Can you deploy this?
Action.Main-->>Commit.Main: Deploying now.
create participant Google.Cloud
Action.Main->>Google.Cloud: Deploy Main
Action.Main->>Google.Cloud: Push Container
Action.Main->>Google.Cloud: Deploy Migrations
Action.Main->>Google.Cloud: Create Deployment
Action.Main-->>Commit.Main: Tag Release
```

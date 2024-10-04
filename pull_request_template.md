## ■ Overview

- Issue: [#Issue number](Link to GitHub issue)
- Redmine: [Ticket number](Link to Redmine issue)

## ■ Design/Cách làm

- Describe shortly what you did
- List subtask (if exists)
    - #600

## ■ Effects/ Phạm vi ảnh hưởng

- Feature/Chức năng:
- Screen/Màn hình:

## ■ Review checklist

- Did you follow the coding standard: https://wiki.zigexn.vn/doku.php?id=coding_stndard
    - [ ]  PSR2/PSR12
    - [ ]  Sonarlint
- Performance: https://wiki.zigexn.vn/doku.php?id=performance_checklist
    - [ ]  Did you minimize the number of requests to the database? Eliminate the N+1 query problem? Are your queries efficient?
    - [ ]  Did you use Caching for a result's method or object if it is called many times/used frequently?
    - [ ]  Put processes that don't need an immediate response into the background processing.
    - [ ]  Have you added indexes to foreign keys, joint tables, and polymorphic relationships?
- Security: https://wiki.zigexn.vn/doku.php?id=security_checklist
    - [ ]  Did you consider what security vulnerabilities this code is susceptible to?
    - [ ]  Are authorization and authentication handled in the right way?
    - [ ]  Is (user) input validated, sanitized, and escaped to prevent security attacks such as cross-site scripting, SQL injection?
    - [ ]  Is sensitive data like user data, credit card information securely handled and stored?
    - [ ]  Does this code change reveal secret information like keys, passwords, or usernames?
    - [ ]  Is data retrieved from external APIs or libraries checked accordingly?
    - [ ]  Does error handling or logging expose us to vulnerabilities?

## ■ Test (Nội dung đã test)

- Unit Test/Đã có Unit Test cover chưa?: Yes/No/Little
- Consider the effects to the older version/Đã cân nhắc ảnh hưởng tới các version trước chưa?: Yes/Not yet
- Testcase/Nội dung test hoặc test case: [Testcase](Link to testcase)
- Evidence: Screenshot || Screen Record || Logs tương ứng với từng test case

## ■ Notes/Điểm lưu ý

- Cần test kỹ các chức năng:

## ■ Release steps/Release note

**Refer:** https://github.com/ZIGExN/dorapita/wiki/GCP-Staging-environment

### Staging Release

***Checkout to feature branch, update repo and deploy***

#### Feature 1

```
# On <instance_name>
# Change the working Directory to "dora-pt.jp"
cd /var/www/dora-pt.jp

# Switch to feature branch
git checkout <feature_branch>
git pull

# Describe command 1
./bin/cake command_1

# Describe command 2
./bin/cake command_2

```

#### Feature 2

```
# On <instance_name>
# Change the working Directory to "legacy.dorapita.com"
cd /var/www/legacy.dorapita.com

# Switch to feature branch
git checkout <feature_branch>
git pull

# Describe command 1
./bin/cake command_1

# Describe command 2
./bin/cake command_2

```

### Production Release

***Merge feature branch into main, update repo and deploy***

#### Feature 1

Merge `<feature_branch_1>` into `main`

```
# Change the working Directory to "dora-pt.jp"
cd /var/www/dora-pt.jp
git pull

# Describe command 1
./bin/cake command_1

# Describe command 2
./bin/cake command_2

```

#### Feature 2

Merge `<feature_branch_2>` into `main`

```
# Change the working Directory to "legacy.dorapita.com"
cd /var/www/legacy.dorapita.com
git pull

# Describe command 1
./bin/cake command_1

# Describe command 2
./bin/cake command_2

```

## ■ Revert steps
### For both Staging and Production
#### 1. Revert database
For minor changes, use `rollback` command to revert:
```
# On <instance_name>
cd /var/www/cadm.dorapita.com

# Check migrations status before rollback
./bin/cake migrations status

# Rollback migrations
./bin/cake migrations rollback -d <previous_version>

# Check migrations status after rollback
./bin/cake migrations status
```

For large data:
Perform a backup of the data/tables before making significant changes.

#### 2. Revert command
```
# On <instance_name>
cd /var/www/legacy.dorapita.com

# Use the --revert option to undo changes related to the command
./bin/cake command_name --revert

# Use this command to update env to the previous value 
.bin/cake replace_env_example
```

#### 3. Revert code (for Production only)
Use GitHub to revert. See attached file:

<img width="798" alt="image" src="https://github.com/user-attachments/assets/c9e1faca-43c8-4fde-8a3b-b827a5eda1ba">

Apply for these pull requests below:

1. Revert <PR_1>
2. Revert <PR_2>

## ■ Evidence for Revert results
### Revert database
#### 1. Check migrations status before rollback
<evidence_photo>
#### 2. Run rollback command
<evidence_photo>
#### 3. Check migrations status after rollback
<evidence_photo>

### Revert command
**Revert command 1**
#### 1. Check before running `revert command 1`
<evidence_photo>
#### 2. Run `revert command 1`
<evidence_photo>
#### 3. Check after running `revert command 1`
<evidence_photo>

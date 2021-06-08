# Terraform_Zoom_Provider

This providre allows to perform Create ,Read ,Update and Delete user for zoom channel

## Requirements

terraform zoom provider is based on Terraform, this means that you need

1. Go lang >=1.11 <br>
2. Terraform>=v0.13.0 <br/>
3. Zoom Pro account (Token)

## Installation

1. Download required binaries<br>
2. move binary ~/.terraform.d/plugins/[architecture name]/

## Comands to Run the Provider

1. terraform init <br/>
2. terraform plan <br>
3. terrafrom apply (To create or update the user)<br>
4. terraform destroy (To destroy the created user)<br>

## Steps to run import command

1. Write manually a resource configuration block for the resource, to which the imported object will be mapped.
2. RUN terraform import zoom_user.sample <user_id> 


## Create User test
 1.For new user test success
 2. Already user user exist confimation
 3. Validation of email id
11:34
GetItem :
User Exist
 2. User doen't exist
11:35
Update:
 1. update if user exist
 2. reject if user doen't exist
11:36
Delete:
1. delete if user already exist
2. reject if user doesn't exist




# Resource Testing

1. On providing wrong data in TestAccItem_Basic(CREATE,READ,DELETE), test must fails.
2. On passing incorrect data wrt  existing user in TestAccItem_Update(UPDATE), test must fails.


# Client testing
1. Create the user and validates its attributes.
2. In NewItem testing,it checks if user exist .
3. In UpdateItem testing ,it updates the user and then matches its attributes and essentialy verifies updates are successful.
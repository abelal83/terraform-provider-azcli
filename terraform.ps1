#go build -o "terraform-provider-cosmos.exe"
#.\terraform-provider-cosmos.exe
Set-Location -Path $PSScriptRoot
Remove-Item -Path $PSScriptRoot\.terraform -Recurse -Force
Remove-Item -Path .\terraform.tfstate
terraform.exe init
terraform.exe plan
terraform.exe apply -auto-approve
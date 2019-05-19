Set-Location $PSScriptRoot
go build -o "terraform-provider-azcli.exe"
Remove-Item .\.terraform -Force -Confirm:$false
terraform init
#.\terraform-provider-cosmos.exe
#terraform.exe init
#terraform.exe plan
#terraform.exe apply -auto-approve
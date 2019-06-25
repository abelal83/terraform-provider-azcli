Set-EnvironmentVariable -Name GOARCH -Value amd64 -ForProcess
Set-Location $PSScriptRoot
Set-EnvironmentVariable -Name GOOS -Value windows -ForProcess
go build -o "terraform-provider-azcli_v0.0.4_x64.exe"
#Remove-Item .\.terraform -Force -Confirm:$false
#terraform init
#.\terraform-provider-cosmos.exe
#terraform.exe init
#terraform.exe plan
#terraform.exe apply -auto-approve
Set-EnvironmentVariable -Name GOOS -Value linux -ForProcess
go build -o "terraform-provider-azcli_v0.0.4_x64"
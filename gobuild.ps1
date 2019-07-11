$env:GOOS = "windows"
$env:GOARCH = "amd64"
$env:TF_LOG = "TRACE"
remove-item ".\terraform-provider-azcli_v0.0.10_x64.exe" -Force
go build -o "terraform-provider-azcli_v0.0.10_x64.exe"
Remove-Item .\.terraform -Force -Recurse
#remove-item .\terraform.tfstate -Force
#terraform init
#terraform.exe plan 
<#
$env:GOOS = "linux"
go build -o "terraform-provider-azcli_v0.0.7_x64"
#>
rem set the TAG env var from the release dir
for /f %%i in ('dir /B release') do set TAG=%%i

rem copy the storagenode binaries to the installer project
copy release\%TAG%\storagenode_windows_amd64.exe installer\windows\storagenode.exe
copy release\%TAG%\storagenode-updater_windows_amd64.exe installer\windows\storagenode-updater.exe

rem install NuGet packages
nuget install installer\windows\StorjTests\packages.config -o installer\windows\packages

rem build the installer
msbuild installer\windows\windows.sln /t:Build /p:Configuration=Release

rem copy the MSI to the release dir
copy installer\windows\bin\Release\storagenode.msi release\%TAG%\storagenode_windows_amd64.msi

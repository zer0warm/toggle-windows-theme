$key = "Registry::HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize"
$newsys = [int]!(Get-ItemProperty -Path $key).SystemUsesLightTheme
$newapp = [int]!(Get-ItemProperty -Path $key).AppsUseLightTheme
Set-ItemProperty -Path $key -Type DWord -Name SystemUsesLightTheme -Value $newsys
Set-ItemProperty -Path $key -Type DWord -Name AppsUseLightTheme -Value $newapp

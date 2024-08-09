import winreg

with winreg.OpenKey(
        winreg.HKEY_CURRENT_USER,
        r'SOFTWARE\Microsoft\Windows\CurrentVersion\Themes\Personalize',
        access=(winreg.KEY_QUERY_VALUE | winreg.KEY_SET_VALUE)) as key:
    curapps, cursys = \
        winreg.QueryValueEx(key, "AppsUseLightTheme"), \
        winreg.QueryValueEx(key, "SystemUsesLightTheme")
    winreg.SetValueEx(key, "AppsUseLightTheme", 0, curapps[1], not curapps[0])
    winreg.SetValueEx(key, "SystemUsesLightTheme", 0, cursys[1], not cursys[0])

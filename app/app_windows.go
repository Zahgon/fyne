//go:build !ci && !android && !ios && !wasm && !test_web_driver && !tinygo

package app

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
)

const notificationTemplate = `$title = %q
$content = %q
$iconPath = "file:///%s"
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
$template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastImageAndText02)
$toastXml = [xml] $template.GetXml()
$toastXml.GetElementsByTagName("text")[0].AppendChild($toastXml.CreateTextNode($title)) > $null
$toastXml.GetElementsByTagName("text")[1].AppendChild($toastXml.CreateTextNode($content)) > $null
$toastXml.GetElementsByTagName("image")[0].SetAttribute("src", $iconPath) > $null
$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($toastXml.OuterXml)
$toast = [Windows.UI.Notifications.ToastNotification]::new($xml)
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("%s").Show($toast);`

const scheduledNotificationTemplate = `$title = %q
$content = %q
$iconPath = "file:///%s"
$id = %q
$delivery = [DateTimeOffset]::Parse(%q)
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
$template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastImageAndText02)
$toastXml = [xml] $template.GetXml()
$toastXml.GetElementsByTagName("text")[0].AppendChild($toastXml.CreateTextNode($title)) > $null
$toastXml.GetElementsByTagName("text")[1].AppendChild($toastXml.CreateTextNode($content)) > $null
$toastXml.GetElementsByTagName("image")[0].SetAttribute("src", $iconPath) > $null
$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($toastXml.OuterXml)
$scheduled = [Windows.UI.Notifications.ScheduledToastNotification]::new($xml, $delivery)
$scheduled.Tag = $id
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("%s").AddToSchedule($scheduled);`

const cancelScheduledNotificationTemplate = `$id = "%s"
[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
$notifier = [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("%s")
foreach ($s in $notifier.GetScheduledToastNotifications()) {
    if ($s.Tag -eq $id) { $notifier.RemoveFromSchedule($s) }
}`

func (a *fyneApp) OpenURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

var scriptNum = 0

func (a *fyneApp) SendNotification(n *fyne.Notification) { _ = "STUB: not implemented"; return }

func (a *fyneApp) ScheduleNotification(n *fyne.Notification, when time.Time) (*fyne.ScheduledNotification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *fyneApp) CancelScheduledNotification(id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *fyneApp) notificationAppID() string { _ = "STUB: not implemented"; return "" }

func (a *fyneApp) SetSystemTrayMenu(menu *fyne.Menu) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayIcon(icon fyne.Resource) { _ = "STUB: not implemented"; return }

func (a *fyneApp) SetSystemTrayWindow(w fyne.Window) { _ = "STUB: not implemented"; return }

func runScript(name, script string) { _ = "STUB: not implemented"; return }

func watchTheme(s *settings) { _ = "STUB: not implemented"; return }

func (a *fyneApp) registerRepositories() { _ = "STUB: not implemented"; return }

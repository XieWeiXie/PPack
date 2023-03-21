package Pack

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type ToApplication interface {
	FileSystem() error
	ToICON() error
	InfoList() error
	ToApp() error
	ToDmg() error
	Clear() error
}

func NewMacApplication(name, icon string) ToApplication {
	return &macApplication{
		name: name,
		icon: icon,
	}
}

func Do(mac ToApplication) error {
	fmt.Println(fmt.Sprintf("🎁 >> Step 1 FileSystem..."))
	if err := mac.FileSystem(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> FileSystem Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("🎁 >> Step 1 FileSystem Done!"))
	fmt.Println(fmt.Sprintf("🎉 >> Step 2 ICON Create..."))
	if err := mac.ToICON(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> Application ICON Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("🎉 >> Step 2 ICON Done!"))
	fmt.Println(fmt.Sprintf("🎈 >> Step 3 Info.plist..."))
	if err := mac.InfoList(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> Application Info.plist Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("🎈 >> Step 3 Info.plist Done!"))
	fmt.Println(fmt.Sprintf("📢 >> Step 4 Application Create..."))
	if err := mac.ToApp(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> Application Bundle Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("📢 >> Step 4 Application Bundle Done!"))
	fmt.Println(fmt.Sprintf("⌛ >> Step 5 Application DMG Create..."))
	if err := mac.ToDmg(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> Application DMG Create Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("⌛ >> Step 5 Application DMG Done!"))
	fmt.Println(fmt.Sprintf("🛁 >> Step 6 Application Consist Create..."))
	if err := mac.Clear(); err != nil {
		fmt.Println(fmt.Sprintf("💔 >> Application Consist Fail!!!"))
		return err
	}
	fmt.Println(fmt.Sprintf("🛁 >> Step 6 Application Consist Done!"))
	fmt.Println(fmt.Sprintf("💖 >> 💖💖💖💖💖💖"))
	return nil
}

type macApplication struct {
	name string
	icon string

	appPath      string
	macOSPath    string
	resourcePath string
	plistPath    string
}

const (
	infoPlist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>English</string>
	<key>CFBundleDisplayName</key>
	<string>{{.applicationName}}</string>
	<key>CFBundleExecutable</key>
	<string>{{.applicationName}}</string>
	<key>CFBundleIconFile</key>
	<string>app.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.XieWeiXie.{{.applicationName}}</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>{{.applicationName}}</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0.0</string>
	<key>CFBundleVersion</key>
	<string>{{.version}}</string>
	<key>CSResourcesFileMapped</key>
	<true/>
	<key>LSApplicationCategoryType</key>
	<string>public.app-category.developer-tools</string>
	<key>LSMinimumSystemVersion</key>
	<string>10.13</string>
	<key>LSRequiresCarbon</key>
	<true/>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSHumanReadableCopyright</key>
	<string></string>
	<key>NSAppTransportSecurity</key>
	<dict>
		<key>NSExceptionDomains</key>
		<dict>
			<key></key>
			<dict>
				<key>NSExceptionAllowsInsecureHTTPLoads</key>
				<true/>
				<key>NSIncludesSubdomains</key>
				<true/>
			</dict>
		</dict>
	</dict>
</dict>
</plist>`
)
const (
	perm = 0755
)

func (m *macApplication) FileSystem() (err error) {
	app := filepath.Join(fmt.Sprintf("%s.app", m.name))
	_, err = os.Stat(app)
	m.appPath = app
	macOSPath := fmt.Sprintf("%s/Contents/MacOS", app)
	resources := fmt.Sprintf("%s/Contents/Resources", app)
	contents := fmt.Sprintf("%s/Contents", app)
	m.macOSPath = macOSPath
	m.resourcePath = resources
	m.plistPath = contents
	if os.IsNotExist(err) {
		_ = os.MkdirAll(macOSPath, perm)
		_ = os.MkdirAll(resources, perm)
	}
	return nil
}

const (
	pListName = "Info.plist"
)

func (m *macApplication) InfoList() error {
	var buf bytes.Buffer
	tmp := template.New("app")
	_, _ = tmp.Parse(infoPlist)
	_ = tmp.Execute(&buf, map[string]string{
		"applicationName": m.name,
		"version":         fmt.Sprintf("%s", time.Now().Format("20060102.150405")),
	})
	path := fmt.Sprintf("%s/%s", m.plistPath, pListName)
	return os.WriteFile(path, buf.Bytes(), perm)
}

func (m *macApplication) ToApp() (err error) {
	srcBin := fmt.Sprintf("%s/%s", m.macOSPath, m.name)
	dstBin, _ := os.Open(m.name)
	f, _ := os.OpenFile(srcBin, os.O_RDWR|os.O_CREATE, 0755)
	_, _ = io.Copy(f, dstBin)
	return
}

const (
	appICON = "app.icns"
	iconSet = "icons.iconset"
)

func (m *macApplication) ToICON() (err error) {
	path := filepath.Join("./", iconSet)
	_ = os.MkdirAll(path, perm)
	sizes := []int{64, 128, 256, 512}
	for i, size := range sizes {
		nameSize := size
		var suffix string
		if i > 0 {
			nameSize = sizes[i-1]
			suffix = "@2x"
		}
		cmd := exec.Command("sips", "-z", fmt.Sprintf("%d", size), fmt.Sprintf("%d", size), m.icon, "--out", filepath.Join(path, fmt.Sprintf("icon_%dx%d%s.png", nameSize, nameSize, suffix)))
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	cmd := exec.Command("iconutil", "-c", "icns", path, "-o", filepath.Join(m.resourcePath, appICON))
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

func (m *macApplication) ToDmg() (err error) {
	_, err = os.Stat(m.plistPath)
	if os.IsNotExist(err) {
		return err
	}
	dmgPath := filepath.Join("./", fmt.Sprintf("%s.dmg", strings.ToUpper(m.name)))
	_ = os.Remove(dmgPath)
	cmd := exec.Command("hdiutil", "create", "-volname", strings.ToUpper(m.name), "-srcfolder", m.appPath, "-ov", "-format", "UDZO", "-o", dmgPath)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func (m *macApplication) Clear() (err error) {
	path := filepath.Join("./", iconSet)
	err = os.RemoveAll(path)
	err = os.RemoveAll(m.appPath)
	return
}

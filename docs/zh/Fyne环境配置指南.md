# Fyne 开发环境配置指南

> 翻译自：https://docs.fyne.io/started/

## 概述

使用 Fyne 工具包构建跨平台应用程序非常简单，但在开始之前需要安装一些工具。如果你的计算机已经配置好了 Go 开发环境，那么以下步骤可能不是必需的，但我们建议阅读针对你操作系统的提示以防万一。如果后续教程步骤失败，你应该重新检查下面的先决条件。

## 先决条件

Fyne 需要 3 个基本元素：
1. **Go 工具**（至少 1.19 版本）
2. **C 编译器**（用于连接系统图形驱动程序）
3. **系统图形驱动程序**

**重要提示**：这些步骤仅用于开发 - 你的 Fyne 应用程序不需要最终用户进行任何设置或依赖安装！

---

## Windows 系统配置

推荐使用 **MSYS2** 平台在 Windows 上工作。按照以下步骤操作：

### 1. 安装 MSYS2

从 [msys2.org](https://www.msys2.org/) 下载并安装 MSYS2

### 2. 打开正确的终端

**重要**：安装完成后，不要使用自动打开的 MSYS 终端！

从开始菜单打开 **"MSYS2 MinGW 64-bit"**

### 3. 安装必要的工具

执行以下命令（如果询问安装选项，请务必选择 "all"）：

```bash
$ pacman -Syu
$ pacman -S git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go
```

### 4. 配置 PATH 环境变量

#### 在 MSYS2 终端中配置

将 `~/Go/bin` 添加到 `$PATH`，在终端中粘贴以下命令：

```bash
$ echo "export PATH=\$PATH:~/Go/bin" >> ~/.bashrc
```

#### 在 Windows 系统中配置

为了让编译器在其他终端中工作，你需要设置 Windows 的 `%PATH%` 环境变量：

1. 打开"编辑系统环境变量"控制面板
2. 点击"高级"选项卡
3. 点击"环境变量"按钮
4. 在"系统变量"中找到 `Path`
5. 添加 `C:\msys64\mingw64\bin` 到 Path 列表中

---

## macOS 系统配置

### 1. 安装 Go

从 [Go 下载页面](https://go.dev/dl/) 下载 Go 并按照说明安装。

### 2. 安装 Xcode

从 Mac App Store 安装 Xcode。

### 3. 设置 Xcode 命令行工具

打开终端窗口并输入以下命令：

```bash
xcode-select --install
```

### 4. 图形驱动程序

在 macOS 中，图形驱动程序已经预装，无需额外安装。

---

## Linux 系统配置

你需要使用包管理器安装 Go、GCC 和图形库头文件。根据你的发行版选择相应的命令：

### Debian、Ubuntu 和 Raspberry Pi OS

```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev
```

### Fedora

```bash
sudo dnf install golang golang-misc gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel libxkbcommon-devel wayland-devel
```

### Arch Linux

```bash
sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi libxkbcommon
```

### Solus

```bash
sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel libxkbcommon-devel
```

### openSUSE

```bash
sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel libxkbcommon-devel
```

### Void Linux

```bash
sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel libXxf86vm-devel libxkbcommon-devel wayland-devel
```

### Alpine Linux

```bash
sudo apk add go gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev libxkbcommon-dev wayland-dev
```

### NixOS

```bash
nix-shell -p libGL pkg-config xorg.libX11.dev xorg.libXcursor xorg.libXi xorg.libXinerama xorg.libXrandr xorg.libXxf86vm libxkbcommon wayland
```

---

## BSD 系统配置

你需要使用包管理器安装 Go、GCC 和图形库头文件。根据你的 BSD 系统选择相应的命令：

### FreeBSD

```bash
sudo pkg install go gcc xorg pkgconf
```

### OpenBSD

```bash
sudo pkg_add go
```

### NetBSD

```bash
sudo pkgin install go pkgconf
```

---

## Android 开发配置

### 方式一：从桌面计算机编译到 Android

1. 首先需要为你当前的计算机（Windows、macOS 或 Linux）安装工具
2. 安装 Android SDK 和 Android NDK
   - **推荐方式**：安装 Android Studio，然后转到 Tools > SDK Manager，从 SDK Tools 安装 NDK (Side by side) 包
   - **精简方式**：下载独立的 Android NDK，解压文件夹并将 `ANDROID_NDK_HOME` 环境变量指向它

### 方式二：在 Android 设备上使用 Termux 编译

[视频教程](https://www.youtube.com/watch?v=XXX)

1. 安装 F-Droid，然后从那里安装 Termux
2. 打开 Termux 并安装 Go 和 Git：
   ```bash
   pkg install golang git
   ```
3. 从 https://github.com/Lzhiyong/termux-ndk 安装 NDK 和 SDK 到 Termux，并适当设置环境变量 `ANDROID_HOME` 和 `ANDROID_NDK_HOME`

---

## iOS 开发配置

1. 要开发 iOS 应用程序，你需要访问 Apple Mac 计算机，并按照上面的 macOS 选项卡进行配置
2. 你还需要创建 Apple Developer 账户并注册开发者计划（需要付费）以获得必要的证书，才能在任何设备上运行你的应用程序

---

## 下载 Fyne

### 1. 创建 Go 模块

从 Go 1.16 开始，你需要在使用包之前设置模块。

运行以下命令并将 `MODULE_NAME` 替换为你喜欢的模块名称（应该在应用程序的新文件夹中调用）：

```bash
$ mkdir myapp
$ cd myapp
$ go mod init MODULE_NAME
```

**提示**：如果你不确定 Go 模块如何工作，可以阅读 [教程：创建 Go 模块](https://go.dev/doc/tutorial/create-module)。

### 2. 下载 Fyne 模块和辅助工具

使用以下命令：

```bash
$ go get fyne.io/fyne/v2@latest
$ go install fyne.io/tools/cmd/fyne@latest
```

---

## 检查安装

在编写应用程序或运行示例之前，你可以使用 **Fyne Setup 工具**检查安装。

只需从链接下载适合你计算机的应用程序并运行它，你应该会看到类似以下的屏幕：

![Fyne Setup Tool](https://docs.fyne.io/started/img/setup.png)

如果安装有任何问题，请参阅故障排除部分以获取提示。

---

## 运行演示程序

如果你想在开始编写自己的应用程序之前看到 Fyne 工具包的实际效果，可以通过执行以下命令在计算机上运行我们的演示应用程序：

```bash
$ go run fyne.io/demo@latest
```

**注意**：第一次运行需要编译一些 C 代码，因此可能比平时花费更长的时间。后续构建会重用缓存，速度会快得多。

### 安装演示程序

如果你愿意，也可以使用以下命令安装演示程序（需要 Go 1.16 或更高版本）：

```bash
$ go install fyne.io/demo@latest
```

如果你的 `GOBIN` 环境变量已添加到 path（在 macOS 和 Windows 上应该是默认的），你可以直接运行演示程序：

```bash
$ demo
```

---

## 完成！

就是这样！现在你可以在你选择的 IDE 中编写自己的 Fyne 应用程序了。

---

## 针对 KrillinAI 项目的说明

### 为什么桌面版无法运行？

如果你的电脑上无法运行 KrillinAI 桌面版，可能是以下原因：

1. **缺少 C 编译器**：Fyne 需要 C 编译器来编译图形驱动程序接口
2. **缺少图形库**：系统缺少必要的图形库头文件
3. **Go 版本过低**：需要 Go 1.19 或更高版本

### 解决方案

#### Windows 用户

1. 按照上面的 Windows 配置步骤安装 MSYS2 和相关工具
2. 确保 `C:\msys64\mingw64\bin` 已添加到系统 PATH
3. 重新编译桌面版：
   ```bash
   cd d:\Sources\KrillinAI
   go build -o KrillinAI_desktop.exe cmd/desktop/main.go
   ```

#### macOS 用户

1. 确保已安装 Xcode 命令行工具
2. 重新编译桌面版：
   ```bash
   cd /path/to/KrillinAI
   go build -o KrillinAI_desktop cmd/desktop/main.go
   ```

#### Linux 用户

1. 按照上面的 Linux 配置步骤安装必要的包
2. 重新编译桌面版：
   ```bash
   cd /path/to/KrillinAI
   go build -o KrillinAI_desktop cmd/desktop/main.go
   ```

### 替代方案：使用 Web 界面

如果配置桌面环境遇到困难，你可以使用 Web 界面版本：

1. 编译服务器版本：
   ```bash
   cd d:\Sources\KrillinAI
   go build -o KrillinAI.exe cmd/server/main.go
   ```

2. 配置 `config/config.toml` 文件

3. 运行服务器：
   ```bash
   ./KrillinAI.exe
   ```

4. 在浏览器中访问：`http://127.0.0.1:8888`

Web 界面提供与桌面版相同的功能，只是运行在浏览器中。

---

## 常见问题排查

### Windows 上的常见问题

**问题**：找不到 gcc 命令

**解决方案**：
- 确保已安装 MSYS2 MinGW 工具链
- 检查 `C:\msys64\mingw64\bin` 是否在系统 PATH 中
- 重启终端或 IDE

**问题**：编译时出现 "undefined reference" 错误

**解决方案**：
- 确保使用 "MSYS2 MinGW 64-bit" 终端，而不是普通的 MSYS2 终端
- 重新安装 mingw-w64-x86_64-toolchain

### macOS 上的常见问题

**问题**：找不到 xcode-select 命令

**解决方案**：
- 从 App Store 安装 Xcode
- 运行 `xcode-select --install`

**问题**：编译时出现权限错误

**解决方案**：
- 使用 `sudo` 运行 xcode-select 命令
- 检查 Xcode 许可协议是否已接受：`sudo xcodebuild -license`

### Linux 上的常见问题

**问题**：找不到某些库文件

**解决方案**：
- 确保安装了所有必需的开发包（-dev 或 -devel 后缀）
- 更新包管理器缓存：`sudo apt update` 或相应的命令
- 检查是否安装了 pkg-config：`which pkg-config`

**问题**：运行时出现 "cannot open display" 错误

**解决方案**：
- 确保 X11 服务器正在运行
- 检查 DISPLAY 环境变量：`echo $DISPLAY`
- 如果使用 SSH，确保启用了 X11 转发

---

## 参考资源

- [Fyne 官方文档](https://docs.fyne.io/)
- [Fyne GitHub 仓库](https://github.com/fyne-io/fyne)
- [Go 官方文档](https://go.dev/doc/)
- [MSYS2 官网](https://www.msys2.org/)
- [KrillinAI 项目文档](./README.md)
- [前端界面说明](./前端界面说明.md)

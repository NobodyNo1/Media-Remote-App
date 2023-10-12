package dev.bekarys.remote_media_control

import platform.UIKit.UIDevice

class IOSPlatform: Platform {
    override val name: String = UIDevice.currentDevice.systemName() + " " + UIDevice.currentDevice.systemVersion

    override fun log(message: String) {
        //TODO: LOG
    }
}

actual fun getPlatform(): Platform = IOSPlatform()
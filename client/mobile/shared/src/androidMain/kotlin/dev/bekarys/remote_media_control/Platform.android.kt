package dev.bekarys.remote_media_control

import android.util.Log

class AndroidPlatform : Platform {
    override val name: String = "Android ${android.os.Build.VERSION.SDK_INT}"

    override fun log(message: String) {
        Log.d("PIPI", message)
    }
}

actual fun getPlatform(): Platform = AndroidPlatform()
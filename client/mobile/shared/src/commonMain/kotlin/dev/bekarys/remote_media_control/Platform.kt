package dev.bekarys.remote_media_control

interface Platform {
    val name: String

    fun log(message: String)
}

expect fun getPlatform(): Platform
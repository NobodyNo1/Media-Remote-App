package dev.bekarys.remote_media_control.domain

import dev.bekarys.remote_media_control.data.MediaControlRepository

class RemoteMediaController {

    //todo: enum
    suspend fun updateVolume(shouldIncrease: Boolean) {
        MediaControlRepository.volume(shouldIncrease)
    }


    suspend fun playPause() {
        MediaControlRepository.playPause()
    }
}
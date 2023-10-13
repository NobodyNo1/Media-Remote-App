package dev.bekarys.remote_media_control.android.screens.home

import androidx.lifecycle.ViewModel
import dev.bekarys.remote_media_control.domain.RemoteMediaController
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.cancel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch
import java.io.Closeable
import kotlin.coroutines.CoroutineContext

data class HomeUiState(
    val loading : Boolean   = false,
    val fail    : String?   = null
)

class CloseableCoroutineScope(
    context: CoroutineContext = SupervisorJob() + Dispatchers.Main.immediate
) : Closeable, CoroutineScope {
    override val coroutineContext: CoroutineContext = context
    override fun close() {
        coroutineContext.cancel()
    }
}

class HomeViewModel(
    private val controller      :RemoteMediaController   = RemoteMediaController(),
    private val coroutineScope  :CloseableCoroutineScope = CloseableCoroutineScope()
): ViewModel(coroutineScope) {

    private val _uiState:   MutableStateFlow<HomeUiState> = MutableStateFlow(HomeUiState())
    private fun MutableStateFlow<HomeUiState>.loading(isLoading: Boolean) {
        val failValue = if(isLoading) null else value.fail
        value = value.copy(loading = isLoading, fail = failValue)
    }
    private fun MutableStateFlow<HomeUiState>.failure(failureText: String?) {
        value = value.copy(fail = failureText)
    }

    fun updateVolume(increase: Boolean) {
        _uiState.loading(true)
        coroutineScope.launch {
            kotlin.runCatching {
                controller.updateVolume(increase)
            }.onFailure {
                _uiState.failure(it.message)
            }
            _uiState.loading(false)
        }
    }

    fun playPause() {
        _uiState.loading(true)
        coroutineScope.launch {
            kotlin.runCatching {
                controller.playPause()
            }.onFailure {
                _uiState.failure(it.message)
            }
            _uiState.loading(false)
        }
    }

    val uiState : StateFlow<HomeUiState> = _uiState

}
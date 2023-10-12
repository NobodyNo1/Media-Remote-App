package dev.bekarys.remote_media_control.android

import android.os.Bundle
import android.widget.Toast
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.material.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import dev.bekarys.remote_media_control.Greeting
import dev.bekarys.remote_media_control.domain.RemoteMediaController
import kotlinx.coroutines.MainScope
import kotlinx.coroutines.launch

class MainActivity : ComponentActivity() {

    private val mainScope = MainScope()
    private val controller = RemoteMediaController()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            MyApplicationTheme {
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colors.background
                ) {
                    ControlView(Greeting().greet())
                }
            }
        }
    }

    @Composable
    fun ControlView(
        text: String,
    ) {
        val isLoading = remember {
            mutableStateOf(false)
        }
        val failState = remember {
            mutableStateOf(Pair(false, ""))
        }
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(),
            verticalArrangement = Arrangement.Center
        ) {
            if(failState.value.first){
                //show error
//            failState.value = Pair(false, "")
                Text("FAILED: ${failState.value.second}", color = Color.Red)
            }
            Spacer(modifier = Modifier.height(8.dp))
            Text(text = text)
            if (isLoading.value)
                CircularProgressIndicator()
            Spacer(modifier = Modifier.height(8.dp))
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceEvenly
            ) {
                MyTextButton("Decrease") {
                    updateVolume(false, {
                        isLoading.value = it
                    }, {
                        isLoading.value = false
                        failState.value = Pair(true, it)
                    })
                }
                MyTextButton("Increase") {
                    updateVolume(true, {
                        isLoading.value = it
                    }, {
                        isLoading.value = false
                        failState.value = Pair(true, it)
                    })
                }
            }
            Spacer(modifier = Modifier.height(8.dp))
            MyTextButton(text = "Play/Pause") {
                playPause({
                    isLoading.value = it
                }, {
                    isLoading.value = false
                    failState.value = Pair(true, it)
                })
            }
        }
    }

    private fun playPause(onStateChange: (Boolean) -> Unit, onFail: (String) -> Unit) {
        mainScope.launch {
            kotlin.runCatching {
                controller.playPause()
            }.onSuccess {
                onStateChange(false)
            }.onFailure {
                onFail("Failed: ${it.localizedMessage}")
            }
            onStateChange(true)
        }
    }

    // TODO: Show state
    private fun updateVolume(
        isIncrease: Boolean,
        onStateChange: (Boolean) -> Unit,
        onFail: (String) -> Unit
    ) {
        mainScope.launch {
            kotlin.runCatching {
                controller.updateVolume(isIncrease)
            }.onSuccess {
                onStateChange(false)
            }.onFailure {
                onFail("Failed: ${it.localizedMessage}")
            }
            onStateChange(true)
        }
    }

    @Composable
    fun MyTextButton(text: String, onClick: () -> Unit) {
        Button(onClick = onClick) {
            Text(text)
        }
    }

    @Preview
    @Composable
    fun DefaultPreview() {
        MyApplicationTheme {
            ControlView("Preview")
        }
    }

}

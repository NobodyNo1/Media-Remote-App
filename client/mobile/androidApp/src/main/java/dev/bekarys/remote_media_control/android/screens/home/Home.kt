package dev.bekarys.remote_media_control.android.screens.home

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.material.Button
import androidx.compose.material.CircularProgressIndicator
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel


@Composable
fun HomeScreen(
    homeViewModel: HomeViewModel = viewModel()
) {
    val uiState by homeViewModel.uiState.collectAsState()
    ControlView(
        uiState.loading,
        uiState.fail,
        homeViewModel::playPause,
        homeViewModel::updateVolume
    )
}

@Composable
fun ControlView(
    isLoading: Boolean,
    failText: String?,
    playPause: () -> Unit,
    updateVolume: (Boolean) -> Unit
) {
    Column(
        verticalArrangement = Arrangement.Center,
        modifier = Modifier
            .fillMaxWidth()
            .fillMaxHeight(),
    ) {

        CircularProgressIndicator(
            modifier = Modifier
                .alpha(if (isLoading) 1.0f else 0.0f)
                .align(Alignment.CenterHorizontally)
        )

        Spacer(modifier = Modifier.height(26.dp))

        Text(
            modifier = Modifier.align(Alignment.CenterHorizontally),
            text = "Volume Control",
            fontSize = 18.sp
        )
        Spacer(modifier = Modifier.height(8.dp))
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceEvenly
        ) {
            MyTextButton(text = "Decrease", enabled = !isLoading) {
                updateVolume(false)
            }
            MyTextButton(text = "Increase", enabled = !isLoading) {
                updateVolume(true)
            }
        }

        Spacer(modifier = Modifier.height(8.dp))

        Text(
            modifier = Modifier.align(Alignment.CenterHorizontally),
            text = "Media Keys",
            fontSize = 18.sp
        )
        Spacer(modifier = Modifier.height(8.dp))
        MyTextButton(
            modifier = Modifier.align(Alignment.CenterHorizontally),
            text = "Play/Pause",
            enabled = !isLoading
        ) {
            playPause()
        }

        failText?.let {
            Spacer(modifier = Modifier.height(8.dp))
            Text(
                modifier = Modifier.align(Alignment.CenterHorizontally),
                text = "FAILED: $it",
                color = Color.Red
            )
            Spacer(modifier = Modifier.height(8.dp))
        }
    }
}

@Composable
fun MyTextButton(
    modifier: Modifier = Modifier,
    text: String,
    enabled: Boolean,
    onClick: () -> Unit
) {
    Button(modifier = modifier, onClick = onClick, enabled = enabled) {
        Text(text)
    }
}

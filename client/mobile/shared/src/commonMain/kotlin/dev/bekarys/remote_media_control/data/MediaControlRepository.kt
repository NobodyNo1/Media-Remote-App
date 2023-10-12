package dev.bekarys.remote_media_control.data

import dev.bekarys.remote_media_control.Platform
import dev.bekarys.remote_media_control.getPlatform
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json

object MediaControlRepository {

    private val platform: Platform = getPlatform()
    //todo: try to convert to gRPC
    private val httpClient = HttpClient {
        install(ContentNegotiation) {
            json(Json {
                ignoreUnknownKeys = true
                useAlternativeNames = false
            })
        }
    }

    suspend fun volume(shouldIncrease: Boolean) {
        val command = "volume"
        val queryEvent = "status"
        val queryValue = if (shouldIncrease) "increase" else "decrease"
        sendCommand(
            command,
            mapOf(queryEvent to queryValue)
        )
    }

    suspend fun playPause() {
        val command = "key"
        val queryEvent = "event"
        val queryValue = "playPause"
        sendCommand(
            command,
            mapOf(queryEvent to queryValue)
        )
    }


    suspend fun sendCommand(
        command: String,
        query: Map<String, String>
    ) {
        val ipAddress = "192.168.0.231"
        val port = "8000"
        val stringBuilder = StringBuilder()
        stringBuilder.append("http://")
        stringBuilder.append(ipAddress).append(':').append(port).append("/")
        stringBuilder.append(command).append("?")
        var isFirst = true
        query.forEach {
            if (!isFirst) {
                stringBuilder.append('&')
                isFirst = false
            }
            stringBuilder.append(it.key)
            stringBuilder.append('=')
            stringBuilder.append(it.value)
        }
        val toString = stringBuilder.toString()
        platform.log(toString)
        val result = httpClient.get(toString)
        platform.log("${result.status}")
    }
}
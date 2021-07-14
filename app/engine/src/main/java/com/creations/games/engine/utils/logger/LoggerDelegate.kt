package com.creations.games.engine.utils.logger

class LoggerDelegate(
    private val level: LogLevel,
    private val log: (logLevel: LogLevel, tag: String?, message: String, showRed: Boolean) -> Unit,
    private val isRed: Boolean = false
) {
    operator fun invoke(message: String, forceRed: Boolean = false) =
        log(level, null, message, isRed || forceRed)

    operator fun invoke(tag: String, message: String, forceRed: Boolean = false) =
        log(level, tag, message, isRed || forceRed)
}
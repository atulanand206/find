package com.creations.games.engine.utils.logger

//this variable is intentionally left with Pascal case. DO NOT CHANGE
// reason: to provide a static access but still keep it differentiated from passed around variables
val Log = Logger()
val LogEngine = Logger("GameEngine")

class Logger(private val identifier: String? = null) {
    var minLogLevel = LogLevel.INFO

    val t = LoggerDelegate(LogLevel.TRACE, ::log)
    val d = LoggerDelegate(LogLevel.DEBUG, ::log)
    val i = LoggerDelegate(LogLevel.INFO, ::log)
    val w = LoggerDelegate(LogLevel.WARNING, ::log)
    val e = LoggerDelegate(LogLevel.ERROR, ::log, true)
    val f = LoggerDelegate(LogLevel.FATAL, ::log, true)

    fun log(l: LogLevel, tag: String?, message: String, showRed: Boolean) {
        if (l < minLogLevel) return
        logWriter.write(identifier, tag, message, showRed)
    }
}
package com.creations.games.engine.utils.logger

object Loggers {
    private val loggers = mutableMapOf<String, Logger>()

    operator fun get(identifier: String) = loggers[identifier] ?: run {
        val logger = Logger(identifier)
        loggers[identifier] = logger
        logger
    }
}
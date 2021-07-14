package com.creations.games.engine.utils.logger

import com.badlogic.gdx.Gdx
import com.creations.games.engine.values.Values

interface LogWriter {
    fun write(identifier: String?, tag: String?, message: String, isRed: Boolean)
}

internal val logWriter: LogWriter get() = when (Gdx.app) {
    null -> ConsoleLogWriter
    else -> ApplicationLogWriter
}

internal object ConsoleLogWriter: LogWriter {
    override fun write(identifier: String?, tag: String?, message: String, isRed: Boolean) {
        val fn = if(isRed) System.err else System.out
        fn.println(craftMessage(message, Values.gameTag, identifier, tag))
    }
}

internal object ApplicationLogWriter: LogWriter {
    override fun write(identifier: String?, tag: String?, message: String, isRed: Boolean) {
        val fn: (String, String) -> Unit = if(isRed) Gdx.app::error else Gdx.app::log
        fn(Values.gameTag, craftMessage(message, identifier, tag))
    }
}

private fun craftMessage(message: String, vararg prefixes: String?): String {
    var prefix = ""
    prefixes.forEach { if(it != null) prefix += "[$it] " }
    return "$prefix$message"
}
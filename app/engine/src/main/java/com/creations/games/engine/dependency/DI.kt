package com.creations.games.engine.dependency

class DI {
    companion object {
        const val default = "default"
    }

    /**
     * All Dependencies are stored wrt hash generated using [getHash]
     */
    internal val dependencies = mutableMapOf<String, Any>()

    fun <T> add(cl: Class<T>, obj: T, name: String = default) where T: Any {
        dependencies[getHash(cl, name)] = obj
    }

    @Suppress("UNCHECKED_CAST")
    operator fun <T> get(cl: Class<T>, name: String = default): T = with(getHash(cl, name)) {
        return@with when(contains(this)) {
            true -> dependencies[this] as T
            else -> error("Cannot find dependency $this")
        }
    }

    fun copyFrom(other: DI) = other.dependencies.forEach { (k, v) -> dependencies[k] = v }

    fun <T> contains(cl: Class<T>, name: String = default) = contains(getHash(cl, name))
    private fun contains(hash: String) = dependencies.contains(hash)

    private fun getHash(type: Class<*>, name: String) = "${type.canonicalName}*$name"
}
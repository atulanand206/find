package com.creations.games.engine.asset

/**
 * Data class storing params for loading assets
 * @param path Path to load asset from
 * @param type Type of asset you are trying to load
 */
data class AssetLoadParam(
    val path: String,
    val type: Class<*>
)
package com.creations.games.engine.dependency

import com.creations.games.engine.asset.GameAssetManager
import com.creations.games.engine.utils.Resizable
import com.creations.games.engine.view.RenderManager

fun DI.resize(w: Int, h: Int) = dependencies.forEach { (_, dep) -> if (dep is Resizable) dep.resize(w, h) }

val DI.renderManager get() = this[RenderManager::class.java]
val DI.assetManager get() = this[GameAssetManager::class.java]
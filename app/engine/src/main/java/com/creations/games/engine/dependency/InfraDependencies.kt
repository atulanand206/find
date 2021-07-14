package com.creations.games.engine.dependency

import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.creations.games.engine.asset.GameAssetManager
import com.creations.games.engine.utils.FileHandleResolverWrapper
import com.creations.games.engine.view.RenderManager

fun DI.setupInfra(resolver: FileHandleResolver) {
    add(RenderManager::class.java, RenderManager())
    add(GameAssetManager::class.java, GameAssetManager(FileHandleResolverWrapper(resolver)))
}
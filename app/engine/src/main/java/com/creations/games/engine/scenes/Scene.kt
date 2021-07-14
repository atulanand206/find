package com.creations.games.engine.scenes

import com.creations.games.engine.dependency.DI
import com.creations.games.engine.dependency.renderManager
import com.creations.games.engine.gameObject.GameObject

abstract class Scene(private val di: DI) {
    //Game objects are updated and drawn in the same order as they are added
    protected val stage get() = di.renderManager.stage

    fun addObjToScene(gameObject: GameObject) {
        di.renderManager.stage.addActor(gameObject)
    }

    abstract fun act(dt:Float)
}
package com.creations.games.engine.scenes

import com.badlogic.gdx.math.Vector2
import com.creations.games.engine.dependency.DI
import com.creations.games.engine.dependency.renderManager
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values

abstract class Scene(private val di: DI) {
    //Game objects are updated and drawn in the same order as they are added
    protected val stage get() = di.renderManager.stage
    protected val screenCenter = Vector2(Values.VIRTUAL_WIDTH/2f, Values.VIRTUAL_HEIGHT/2f)

    fun background(gameObject: GameObject) {
        addObjToScene(gameObject)
    }

    fun addObjToScene(gameObject: GameObject) {
        di.renderManager.stage.addActor(gameObject)
    }

    abstract fun act(dt:Float)
}


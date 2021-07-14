package com.creations.games.engine.view

import com.badlogic.gdx.Gdx
import com.badlogic.gdx.graphics.GL20
import com.badlogic.gdx.graphics.OrthographicCamera
import com.badlogic.gdx.scenes.scene2d.Stage
import com.badlogic.gdx.utils.Disposable
import com.badlogic.gdx.utils.viewport.ExtendViewport
import com.creations.games.engine.utils.Resizable
import com.creations.games.engine.values.Values

class RenderManager: Disposable, Resizable {
    val cam: OrthographicCamera = OrthographicCamera(
        Values.VIRTUAL_WIDTH,
        Values.VIRTUAL_HEIGHT
    ).apply {
        setToOrtho(false, Values.VIRTUAL_WIDTH, Values.VIRTUAL_HEIGHT)
    }

    val vp: ExtendViewport = ExtendViewport(
        Values.VIRTUAL_WIDTH,
        Values.VIRTUAL_HEIGHT,
        cam
    )

    val stage: Stage = Stage(vp).apply {
        batch.enableBlending()
        batch.setBlendFunction(GL20.GL_SRC_ALPHA, GL20.GL_ONE_MINUS_SRC_ALPHA)
    }

    init {
        cam.position.set(vp.worldWidth / 2f, vp.worldHeight / 2f, 0f)
        Gdx.input.inputProcessor = stage
    }

    fun update(dt: Float) = stage.act(dt)
    fun draw() = stage.draw()

    override fun resize(width: Int, height: Int) = vp.update(width, height)

    override fun dispose() {
        stage.dispose()
    }
}
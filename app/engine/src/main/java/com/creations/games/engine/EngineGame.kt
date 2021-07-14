package com.creations.games.engine

import com.badlogic.gdx.ApplicationAdapter
import com.badlogic.gdx.Gdx
import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.badlogic.gdx.graphics.GL20
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.dependency.*
import com.creations.games.engine.values.Values

abstract class EngineGame(
    private val resolver: FileHandleResolver,
    defaultAspectRatio: Pair<Float, Float> = Pair(9f, 16f)
): ApplicationAdapter() {
    protected val di = DI()

    init {
        Values.aspectRatio = defaultAspectRatio
    }

    /**
     * Libgdx is created in this method. You cannot use libgdx specific things before this
     */
    override fun create() {
        di.setupInfra(resolver)
        onCreate()
        di.assetManager.load()
        onLoaded()
    }

    /**
     * Render is called by libgdx in each frame. Use this as a game loop
     */
    override fun render() {
        Gdx.gl.glClearColor(1f, 1f, 1f, 1f)
        Gdx.gl.glClear(GL20.GL_COLOR_BUFFER_BIT)

        val dt = Gdx.graphics.deltaTime

        //update
        di.renderManager.update(dt)
        update(dt)

        val batch = di.renderManager.stage.batch

        //draw
        di.renderManager.draw()
        batch.begin()
        draw(batch)
        batch.end()
    }

    /**
     * On Create is called when libgdx is created.
     * You can only use libgdx components after this call
     *
     * Use this method to add your assets to load
     * After this, all assets that are queued to load will be loaded
     */
    abstract fun onCreate()

    /**
     * Called once all queued assets are loaded.
     * You can create your game here and use the assets you loaded
     *
     * Note that resizing is not completed here.
     * If you need to use resized screen, create an initIfNeededMethod from your update loop
     */
    abstract fun onLoaded()

    /**
     * Update is called once every frame
     * This is called before [draw] is called
     * Use this to update state of your game
     *
     * @param dt time elapsed since last update
     */
    abstract fun update(dt: Float)

    /**
     * Draw is called once every frame
     * This is called after [update] is called
     * Use this for drawing your objects
     *
     * @param batch batch you can use to draw objects
     */
    abstract fun draw(batch: Batch)

    override fun resize(width: Int, height: Int) = di.resize(width, height)
}
package com.creations.games.think

import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.badlogic.gdx.graphics.g2d.Batch
import com.creations.games.engine.EngineGame
import com.creations.games.engine.dependency.assetManager
import com.creations.games.engine.scenes.Scene
import com.creations.games.engine.values.Values
import com.creations.games.think.asset.Assets
import com.creations.games.think.scenes.launcher.SceneLauncher
import com.creations.games.think.scenes.main.MainLauncher
import java.util.*
import kotlin.concurrent.schedule

class Play(
    resolver: FileHandleResolver
) : EngineGame(resolver, defaultAspectRatio = Pair(9f, 16f)) {
    private lateinit var scene:Scene

    init{
        //todo.note - Change screen scale below
        Values.screenFactor = 40f;
    }

    //This is called at the beginning. We load the assets first
    override fun onCreate() {
        //add assets to DI(Dependency Injector)
        di.add(Assets::class.java, Assets(di.assetManager))
    }

    //At this point assets are loaded. You should create your game here
    override fun onLoaded() {
//        setScene(SceneLauncher(di))

//        Timer().schedule(2000) {
            setScene(MainLauncher(di))
//        }
    }

    // adds a scene to the injector
    private fun setScene(scene: Scene) {
        this.scene = scene
        di.add(Scene::class.java, scene)
    }

    //update is called once every frame
    override fun update(dt: Float) {
        scene.act(dt)
    }

    override fun draw(batch: Batch) {

    }
}
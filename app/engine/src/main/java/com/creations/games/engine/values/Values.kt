package com.creations.games.engine.values

import com.badlogic.gdx.Gdx
import com.badlogic.gdx.math.Vector2
import com.creations.games.engine.utils.AspectRatioCalculator

object Values {
    var gameTag = "StarterKit"
    /**
     * Aspect Ratio Multiplier
     * <p>
     * Changing this will result in different screen size, while keeping the aspect ratio
     */
    var screenFactor = 40f
        set(value) {
            require(value >= 0)
            field = value
        }

    /**
     * Aspect Ratio at which game is initially developed
     */
    var aspectRatio = Pair(9f, 16f)
        set(value) {
            require(value.first >= 0 && value.second >= 0)
            field = value
        }

    val actualAspectRatio: Vector2 by lazy {
        AspectRatioCalculator.transform(Gdx.graphics.width.toFloat(), Gdx.graphics.height.toFloat())
    }

    val isPortrait: Boolean by lazy { aspectRatio.second > aspectRatio.first }

    /**
     * Factor by which font is scaled up for loading
     * <p>
     * Higher the value, more loading time wll be needed
     */
    var fontScaling = 4
        set(value) {
            require(value > 0)
            field = value
        }

    val VIRTUAL_WIDTH get() = aspectRatio.first * screenFactor
    val VIRTUAL_HEIGHT get() = aspectRatio.second * screenFactor

    val ACTUAL_WIDTH get() = actualAspectRatio.x * screenFactor
    val ACTUAL_HEIGHT get() = actualAspectRatio.y * screenFactor
}
package com.creations.games.engine.utils

import com.badlogic.gdx.math.Vector2
import com.badlogic.gdx.utils.viewport.Viewport

/**
 * Utility to calculate aspect ratio
 */
object AspectRatioCalculator {
    /**
     * Calculates aspect ratio
     * @param viewport viewport to calculate aspect ratio for
     */
    fun calculate(viewport: Viewport): Vector2 {
        return transform(viewport.worldWidth, viewport.worldHeight)
    }

    fun transform(x : Float, y : Float) : Vector2  = transform(Vector2(x, y))
    private fun transform(v: Vector2): Vector2 {
        v.nor().scl(9f / v.x)
        return v
    }
}
package com.creations.games.engine.asset

import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.graphics.g2d.BitmapFont
import com.badlogic.gdx.graphics.g2d.freetype.FreeTypeFontGenerator
import com.creations.games.engine.utils.FileHandleResolverWrapper
import java.util.*

/**
 * Data class to store font load params
 *
 * @param path font file path
 * @param fontSize size to load at
 *
 * @property font loaded font. Null, if loading is not done
 */
data class TtfFontLoadParam(val path: String, val fontSize: Int) {
    var font: BitmapFont? = null

    /**
     * Loads the font based on [path] and [fontSize] into [font]
     *
     * @param resolver File Handle Resolver to use to use to resolve path
     */
    fun load(resolver: FileHandleResolver) {
        val param = FreeTypeFontGenerator.FreeTypeFontParameter().apply {
            size = fontSize
            color = Color.WHITE
            minFilter = Texture.TextureFilter.Linear
            magFilter = Texture.TextureFilter.Linear
        }

        val generator = FreeTypeFontGenerator(resolver.resolve(getModifiedPath())).apply {
            scaleForPixelHeight(fontSize)
        }

        val fnt = generator.generateFont(param)
        font = fnt
    }

    private fun getModifiedPath(): String {
        return getModifiedPath(processPath(path), fontSize)
    }

    companion object {
        internal const val suffix = ".ttf"

        /**
         * Modifies the path to load font at.
         * This is needed to load font at multiple sizes
         * <p>
         * libgdx Asset Manager does not allow loading of same file in different way.
         * When same font is being asked to load at multiple size, asset manager assumes that it has already been loaded.
         * To bypass this, we use our [FileHandleResolverWrapper] to trick asset manager, such that asset manager thinks they are different files
         *
         * @see [FileHandleResolverWrapper]
         *
         * @param path path to file
         * @param size load size
         *
         * @return resulting path
         */
        fun getModifiedPath(path: String, size: Int): String {
            return "${FileHandleResolverWrapper.delimiter}${processPath(path)}${FileHandleResolverWrapper.delimiter}$size${FileHandleResolverWrapper.delimiter}$suffix"
        }

        /**
         * Removes suffix from path
         *
         * @param string unprocessed path
         * @return processed path without suffix
         */
        fun processPath(string: String): String {
            return if (string.substring(string.length - suffix.length)
                    .toLowerCase(Locale.getDefault()) == suffix
            ) {
                string.substring(0, string.length - suffix.length)
            } else string
        }
    }
}
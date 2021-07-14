package com.creations.games.engine.asset

import com.badlogic.gdx.assets.AssetManager
import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.badlogic.gdx.graphics.Texture
import com.badlogic.gdx.graphics.g2d.BitmapFont
import com.badlogic.gdx.graphics.g2d.freetype.FreeTypeFontGenerator
import com.badlogic.gdx.graphics.g2d.freetype.FreeTypeFontGeneratorLoader
import com.badlogic.gdx.graphics.g2d.freetype.FreetypeFontLoader
import com.badlogic.gdx.utils.Disposable
import com.creations.games.engine.utils.FileHandleResolverWrapper
import com.creations.games.engine.utils.logger.LogEngine
import com.creations.games.engine.values.Values

/**
 * Class to load assets using the provided [fileHandleResolver]
 * NOTE: Doesn't support loading after load() ss called
 */
class GameAssetManager(private val fileHandleResolver: FileHandleResolver) : Disposable {
    private val listeners = mutableListOf<AssetLoadListener>()
    private var assetsToLoad = mutableListOf<AssetLoadParam>()
    private var manager = AssetManager(fileHandleResolver)

    private var fonts = mutableMapOf<String, TtfFontLoadParam>()
    private var loaded = false

    init {
        setupTtfFontLoader()
    }

    /**
     * Configures asset manager to allow loading for TTF font
     */
    private fun setupTtfFontLoader() {
        val fileHandleResolverWrapper = FileHandleResolverWrapper(fileHandleResolver)

        //to be able to load ttf fonts
        manager.setLoader(
            FreeTypeFontGenerator::class.java,
            FreeTypeFontGeneratorLoader(fileHandleResolverWrapper)
        )
        manager.setLoader(
            BitmapFont::class.java,
            TtfFontLoadParam.suffix,
            FreetypeFontLoader(fileHandleResolverWrapper)
        )
    }

    /**
     * Adds assets to load
     * <p>
     * <b> Be sure to call [load] after adding </b>
     *
     * @param path path to load from. Path is wrt point where the project is run at and is resolved using the provided file handle resolver
     * @param type type to load as
     */
    fun addAssetsToLoad(path: String, type: Class<*>) {
        if (loaded) LogEngine.e("Assets have already been loaded instantly. Are you sure, you are calling this at correct time")
        addAssetsToLoad(AssetLoadParam(path, type))
    }

    /**
     * Adds assets to load
     * <p>
     * <b> Be sure to call [load] after adding </b>
     *
     * @param assetLoadParam [AssetLoadParam] to load
     */
    fun addAssetsToLoad(vararg assetLoadParam: AssetLoadParam) {
        if (loaded) LogEngine.e("Assets have already been loaded instantly. Are you sure, you are calling this at correct time")
        assetLoadParam.forEach { assetsToLoad.add(it) }
    }

    /**
     * Adds font to load
     * <p>
     * <b> Be sure to call [load] after adding </b>
     *
     * @param path path to load from. Path is wrt point where the project is run at and is resolved using the provided file handle resolver
     * @param size size to load font at
     */
    fun addFontToLoad(path: String, size: Int) {
        if (loaded) LogEngine.e("Assets have already been loaded instantly. Are you sure, you are calling this at correct time")
        val fontFileName = TtfFontLoadParam.getModifiedPath(path, size)
        fonts[fontFileName] ?: let {
            fonts[fontFileName] = TtfFontLoadParam(path, (size * Values.fontScaling))
        }
    }

    /**
     * Loads all assets added in [manager]
     */
    fun load() {
        assetsToLoad.forEach { manager.load(it.path, it.type) }
        fonts.forEach { (_, f) -> f.font ?: f.load(fileHandleResolver) }

        manager.finishLoading()
        loaded = true

        listeners.forEach { l -> l.onLoaded() }
    }

    /**
     * Returns a font if it loaded
     * <p>
     * Trying to get an unloaded font will result in error.
     * <b> Fonts at different sizes are to be loaded separately
     *
     * @param path font file path. Path is wrt point where the project is run at and is resolved using the provided file handle resolver
     * @param size font size
     *
     * @return font
     */
    fun getFont(path: String, size: Int): BitmapFont {
        val fontFileName = TtfFontLoadParam.getModifiedPath(path, size)
        fonts[fontFileName] ?: error("font not loaded, path : $path, size : $size")

        return fonts[fontFileName]?.font ?: error("font loading not complete : $path, size ; $size")
    }

    /**
     * Returns an asset if it is loaded
     * <p>
     * Trying to get an unloaded asset will result in error
     *
     * @param path asset load path
     * @param type asset Type
     *
     * @return asset
     */
    fun <T> get(path: String, type: Class<T>): T {
        val asset = manager.get(path, type)
        if(asset is Texture){
            asset.setFilter(Texture.TextureFilter.Linear, Texture.TextureFilter.Linear)
        }
        return asset
    }

    fun addOnLoadListener(listener: AssetLoadListener) {
        listeners.add(listener)
    }

    override fun dispose() = manager.dispose()
}


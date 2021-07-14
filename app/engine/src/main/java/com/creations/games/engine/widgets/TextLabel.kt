package com.creations.games.engine.widgets

import com.badlogic.gdx.graphics.Color
import com.badlogic.gdx.graphics.g2d.GlyphLayout
import com.badlogic.gdx.scenes.scene2d.Actor
import com.badlogic.gdx.scenes.scene2d.ui.Label
import com.badlogic.gdx.scenes.scene2d.ui.Label.LabelStyle
import com.badlogic.gdx.scenes.scene2d.ui.Skin
import com.badlogic.gdx.scenes.scene2d.ui.WidgetGroup
import com.badlogic.gdx.utils.StringBuilder
import com.creations.games.engine.gameObject.GameObject
import com.creations.games.engine.values.Values
import com.creations.games.engine.values.Values.fontScaling

/**
 * Alternative to [Label] with added features
 * <p>
 *     Scale
 *     Rotate
 *     Crisp Font
 *
 * @property padding Padding controls for the label container
 */
open class TextLabel private constructor(private val label: Label) : GameObject() {
    constructor(text: CharSequence?, style: LabelStyle?) : this(Label(text, style))
    constructor(text: CharSequence?, skin: Skin?) : this(Label(text, skin))
    constructor(text: CharSequence?, skin: Skin?, styleName: String?) : this(Label(text, skin, styleName))
    constructor(text: CharSequence?, skin: Skin?, fontName: String?, color: Color?) : this(Label(text, skin, fontName, color))
    constructor(text: CharSequence?, skin: Skin?, fontName: String?, colorName: String?) : this(Label(text, skin, fontName, colorName))

    private val labelScalingContainer = WidgetGroup(label) //label cannot be scaled (issue of libgdx), so adding it to a container that can be scaled
    val padding = Padding()

    init {
        addActor(labelScalingContainer)
        label.setFillParent(true)

        this.setSize(reqWidth, reqHeight)
    }

    override fun layout() {
        super.layout()
        adjustLabel()
    }

    //blocking override of addActor
    final override fun addActor(actor: Actor?) = super.addActor(actor)

    /**
     * adjusts position and size
     */
    private fun adjustLabel() {
        labelScalingContainer.setPosition(padding.left, padding.bottom)
        labelScalingContainer.setSize(
            (width - padding.left - padding.right) * fontScaling,
            (height - padding.top - padding.bottom) * fontScaling
        )
        labelScalingContainer.setScale(1f / fontScaling, 1f / fontScaling)
    }

    //region delegate properties
    var style: LabelStyle
        get() = label.style
        set(value) {
            label.style = value
        }

    val glyphLayout: GlyphLayout get() = label.glyphLayout
    val text: StringBuilder get() = label.text
    val labelAlign get() = label.labelAlign
    val lineAlign get() = label.lineAlign
    //endregion

    //region delegate methods
    fun setText(value: Int) = label.setText(value)
    fun setText(newText: CharSequence) = label.setText(newText)
    fun setAlignment(alignment: Int) = label.setAlignment(alignment)
    fun setAlignment(labelAlign: Int, lineAlign: Int) = label.setAlignment(labelAlign, lineAlign)
    fun setEllipsis(ellipsis: Boolean) = label.setEllipsis(ellipsis)
    fun setEllipsis(ellipsis: String) = label.setEllipsis(ellipsis)
    fun setWrap(wrap: Boolean) = label.setWrap(wrap)
    override fun toString(): String = label.toString()
    val reqWidth get() = label.prefWidth / fontScaling + padding.left + padding.right
    val reqHeight get() = label.prefHeight / Values.fontScaling + padding.top + padding.bottom
    //endregion

    data class Padding(
        var top: Float = 0f,
        var bottom: Float = 0f,
        var left: Float = 0f,
        var right: Float = 0f,
    )
}
package com.creations.games.engine.utils

import com.badlogic.gdx.assets.loaders.FileHandleResolver
import com.badlogic.gdx.files.FileHandle
import com.creations.games.engine.utils.FileHandleResolverWrapper.Companion.delimiter

/**
 * Resolves File handle to support multiple file name for same file
 * <p>
 * Uses [delimiter] to detect if the path is correct or needs to be resolved separately
 * <p>
 * Assumption : Any handle that starts with the delimiter is to be processed
 * The first delimiter will be ignored as it just marks where the handle is to processed
 * After the first delimiter, every text between a pair of delimiters will be ignored
 * If a delimiter (apart from initial one) doesn't have a pair, the file name will be assumed to correct all along and won't be modified
 *
 * @param resolver File Handle resolver already being used
 */
class FileHandleResolverWrapper(private val resolver: FileHandleResolver) :
    FileHandleResolver {
    private var stringBuilder = java.lang.StringBuilder()

    override fun resolve(fileName: String): FileHandle {
        if (fileName.length <= 2) {
            return resolver.resolve(fileName)
        }

        /*
		 * Only process filename if it starts with our delimiter
		 */
        if (fileName.substring(0, 2) != delimiter) return resolver.resolve(fileName)

        var newFileName = ""
        stringBuilder = StringBuilder()
        val strings = fileName.split(delimiter.toRegex()).toTypedArray()

        /*
		 * Number of delimiters will always be odd
		 * This means number of splits should always be even
		 * In case it is not, just send it directly to process
		 */
        if (strings.size % 2 == 1) return resolver.resolve(fileName)

        /*
		 * Add all odd splits
		 */
        var i = 1
        while (i < strings.size) {
            newFileName = stringBuilder.append(strings[i]).toString()
            i += 2
        }
        return resolver.resolve(newFileName)
    }

    companion object {
        var delimiter = "~~"
    }
}
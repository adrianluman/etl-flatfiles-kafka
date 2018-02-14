// package etl.flatfiles

import scala.io.Source
import java.nio.charset.CodingErrorAction
import scala.io.Codec

object Extract extends App {
	implicit val codec = Codec("UTF-8")
	codec.onMalformedInput(CodingErrorAction.IGNORE)
	codec.onUnmappableCharacter(CodingErrorAction.IGNORE)

	val bufferedSource = Source.fromFile("/home/test/goprojects/src/bitbucket.org/datafeedetl/testfile")
	for (line <- bufferedSource.getLines) {
		println(line)
	}

	bufferedSource.close
	// val file = getClass.getResource("/asdf.txt").getFile()

	// // Source.fromFile(file).getLines.foreach(println)
	// val bufferedSource = Source.fromFile(file)
	// Source.fromInputStream(getClass.getResourceAsStream("/asdf.txt")).getLines().foreach(println)

}

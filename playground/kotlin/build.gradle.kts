plugins {
    kotlin("jvm") version "1.9.24"
    kotlin("plugin.serialization") version "1.9.24"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("com.algolia:algoliasearch-client-kotlin")
    implementation("io.ktor:ktor-client-okhttp:2.3.11")
    implementation("ch.qos.logback:logback-classic:1.5.6")
    implementation("io.github.cdimascio:dotenv-kotlin:6.4.1")
}

tasks.test {
    useJUnitPlatform()
}
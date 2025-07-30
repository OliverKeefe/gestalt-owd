package net.gestalt.app.service.password;

import de.mkammerer.argon2.Argon2;
import de.mkammerer.argon2.Argon2Factory;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Singleton;

@Singleton
public class PasswordService {
    private final Argon2 argon2 = Argon2Factory.create();

    public String hashPassword(String password) {
        int iterations = 12;
        int memory = 65536;
        int parallelism = 1;

        return argon2.hash(iterations,memory, parallelism, password.toCharArray());
    }

    public boolean verifyPassword(String passwordHash, String passwordPlainText) {
        return argon2.verify(passwordHash, passwordPlainText.toCharArray());
    }
}

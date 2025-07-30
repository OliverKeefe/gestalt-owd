package net.gestalt.app.service.user;

import jakarta.inject.Inject;
import net.gestalt.api.dto.login.LoginRequest;
import net.gestalt.app.service.password.PasswordService;

public class VerifyCredentials {

    @Inject
    UserService userService;

    @Inject
    PasswordService passwordService;

    public boolean verifyUsername(String username) {
        return userService.findByUsername(username);
    }

    public boolean verifyPassword(String password) {
        return passwordService.verifyPassword(password);
    }
}

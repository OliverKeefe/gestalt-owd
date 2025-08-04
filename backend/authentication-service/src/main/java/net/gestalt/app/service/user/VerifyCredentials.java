package net.gestalt.app.service.user;

import jakarta.inject.Inject;
import net.gestalt.api.dto.login.LoginRequest;
import net.gestalt.app.service.password.PasswordService;
import net.gestalt.data.entity.User;

import java.util.Objects;

public class VerifyCredentials {

    @Inject
    UserService userService;

    @Inject
    PasswordService passwordService;

    public boolean verifyUsername(String username) {
        User findUser = userService.findByUsername(username);
        return Objects.equals(findUser.getUsername(), username);
    }

    public boolean verifyPassword(String password) {
        return passwordService.verifyPassword(password);
    }
}

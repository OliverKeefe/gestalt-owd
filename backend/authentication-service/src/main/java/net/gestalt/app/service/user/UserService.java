package net.gestalt.app.service.user;

import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;
import net.gestalt.api.dto.login.LoginRequest;
import net.gestalt.data.entity.User;
import net.gestalt.data.repository.UserRepository;

@ApplicationScoped
public class UserService {

    @Inject
    UserRepository userRepository;

    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    public
}

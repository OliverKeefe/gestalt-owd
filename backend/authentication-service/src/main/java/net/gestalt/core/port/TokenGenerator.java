package net.gestalt.core.port;

import net.gestalt.core.model.User;

public interface TokenGenerator {
    String generate(User user);
}

package net.gestalt.api.dto.auth;

import java.util.List;

public record AuthResponse(String username, List<String> roles) {}

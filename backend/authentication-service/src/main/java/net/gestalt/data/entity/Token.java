package net.gestalt.data.entity;

import jakarta.persistence.Id;

public class Token {
    @Id
    private Long id;
    private String token;
    private User user;
}

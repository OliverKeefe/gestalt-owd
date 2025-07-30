package net.gestalt.data.entity;

import jakarta.persistence.Id;

import java.util.UUID;

public class RefreshToken {
    @Id
    private UUID id;

    private String issuedAt;

    private String expiresAt;



}

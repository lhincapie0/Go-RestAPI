package com.example.restapi.model;

public class Server {

    public String ipAddress;
    public String grade;
    public String country;
    public String owner;

    public Server() {
    }

    public Server(String ipAddress, String grade, String country, String owner) {
        this.ipAddress = ipAddress;
        this.grade = grade;
        this.country = country;
        this.owner = owner;
    }

    public String getIpAddress() {
        return ipAddress;
    }

    public void setIpAddress(String ipAddress) {
        this.ipAddress = ipAddress;
    }

    public String getGrade() {
        return grade;
    }

    public void setGrade(String grade) {
        this.grade = grade;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }
}

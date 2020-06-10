package com.example.restapi.model;


import java.io.Serializable;
@SuppressWarnings("serial")
public class HistoryElement  implements Serializable {

    private String host;
    private Info info;
    public HistoryElement() {
    }

    public Info getInfo() {
        return info;
    }

    public void setInfo(Info info) {
        this.info = info;
    }

    public HistoryElement(String host, Info info) {
        this.host = host;
        this.info = info;

    }

    public String getHost() {
        return host;
    }

    public void setHost(String host) {
        this.host = host;
    }



}

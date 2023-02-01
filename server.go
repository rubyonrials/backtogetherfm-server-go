package main

import (
    "log"
    "net/http"
    "time"
    "encoding/json"
    "github.com/livekit/protocol/auth"
    "github.com/google/uuid"
)

func issueTokens(w http.ResponseWriter, r *http.Request) {
    tokens := make(map[string]string)
    redToken, err1 := getJoinToken("red");
    blueToken, err2 := getJoinToken("blue");
    greenToken, err3 := getJoinToken("green");

    if err1 == nil && err2 == nil && err3 == nil {
      tokens["red"] = redToken;
      tokens["blue"] = blueToken;
      tokens["green"] = greenToken;
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    json.NewEncoder(w).Encode(tokens)
}

func getJoinToken(room string) (string, error) {
    canPublish := true
    canSubscribe := true

    at := auth.NewAccessToken(process.env.LIVEKIT_API_KEY, process.env.LIVEKIT_API_SECRET)
    grant := &auth.VideoGrant{
        RoomJoin:     true,
        Room:         room,
        CanPublish:   &canPublish,
        CanSubscribe: &canSubscribe,
    }
    at.AddGrant(grant).
        SetIdentity(uuid.New().String()).
        SetValidFor(time.Hour)

    return at.ToJWT()
}

func main() {
    http.HandleFunc("/issue-tokens", issueTokens)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

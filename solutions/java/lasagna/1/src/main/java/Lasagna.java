public class Lasagna {
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int currentTime) {
        return expectedMinutesInOven() - currentTime;
    }

    public int preparationTimeInMinutes(int numLayers) {
        return 2 * numLayers;
    }

    public int totalTimeInMinutes(int numLayers, int currentTime) {
        return preparationTimeInMinutes(numLayers) + currentTime;
    }
}
